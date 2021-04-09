/*
 *SPDX-License-Identifier: Apache-2.0
 */

const path = require('path');
const fs = require('fs-extra');
const sha = require('js-sha256');
const asn = require('asn1.js');
const { Utils } = require('fabric-common');
const FabricClient = require('./../FabricClient.js');
const ExplorerError = require('../../../common/ExplorerError');
const explorer_error = require('../../../common/ExplorerMessage').explorer
	.error;
const helper = require('../../../common/helper');

const logger = helper.getLogger('FabricUtils');

async function createFabricClient(config, persistence) {
	// Create new FabricClient
	const client = new FabricClient(config);
	// Initialize fabric client
	logger.debug(
		'************ Initializing fabric client for [%s]************',
		config.getNetworkId()
	);
	try {
		await client.initialize(persistence);
		return client;
	} catch (err) {
		throw new ExplorerError(explorer_error.ERROR_2014);
	}
}

/**
 *
 *
 * @param {*} dateStr
 * @returns
 */
async function getBlockTimeStamp(dateStr) {
	try {
		return new Date(dateStr);
	} catch (err) {
		logger.error(err);
	}
	return new Date(dateStr);
}

/**
 *
 *
 * @returns
 */
async function generateDir() {
	const tempDir = `/tmp/${new Date().getTime()}`;
	try {
		fs.mkdirSync(tempDir);
	} catch (err) {
		logger.error(err);
	}
	return tempDir;
}

/**
 *
 *
 * @param {*} header
 * @returns
 */
async function generateBlockHash(header) {
	const headerAsn = asn.define('headerAsn', function() {
		this.seq().obj(
			this.key('Number').int(),
			this.key('PreviousHash').octstr(),
			this.key('DataHash').octstr()
		);
	});
	logger.info('generateBlockHash', header.number.toString());
	// ToDo: Need to handle Long data correctly. header.number {"low":3,"high":0,"unsigned":true}
	const output = headerAsn.encode(
		{
			Number: parseInt(header.number.low),
			PreviousHash: header.previous_hash,
			DataHash: header.data_hash
		},
		'der'
	);
	return sha.sha256(output);
}

/**
 *
 *
 * @param {*} config
 * @returns
 */
function getPEMfromConfig(config) {
	let result = null;
	if (config) {
		if (config.path) {
			// Cert value is in a file
			try {
				result = readFileSync(config.path);
				result = Utils.normalizeX509(result);
			} catch (e) {
				logger.error(e);
			}
		}
	}

	return result;
}

/**
 *
 *
 * @param {*} config_path
 * @returns
 */
function readFileSync(config_path) {
	try {
		const config_loc = path.resolve(config_path);
		const data = fs.readFileSync(config_loc);
		return Buffer.from(data).toString();
	} catch (err) {
		logger.error(`NetworkConfig101 - problem reading the PEM file :: ${err}`);
		throw err;
	}
}

exports.generateBlockHash = generateBlockHash;
exports.createFabricClient = createFabricClient;
exports.getBlockTimeStamp = getBlockTimeStamp;
exports.generateDir = generateDir;
exports.getPEMfromConfig = getPEMfromConfig;
