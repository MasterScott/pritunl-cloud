/// <reference path="References.d.ts"/>
import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as Blueprint from '@blueprintjs/core';
import Main from './components/Main';
import * as Event from './Event';
import * as Csrf from './Csrf';

Csrf.load().then((): void => {
	Blueprint.FocusStyleManager.onlyShowFocusOnTabs();
	Event.init();

	ReactDOM.render(
		<div><Main/></div>,
		document.getElementById('app'),
	);
});
