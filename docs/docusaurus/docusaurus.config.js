/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

module.exports = {
  title: 'Terraform-Validator',
  tagline: 'A norms and conventions validator for Terraform',
  url: 'https://thazelart.github.io/',
  baseUrl: '/terraform-validator/',
  favicon: 'img/terraform-validator-icon.svg',
  organizationName: 'thazelart', // Usually your GitHub org/user name.
  projectName: 'terraform-validator', // Usually your repo name.

  themeConfig: {
    image: 'img/terraform-validator.png',
    disableDarkMode: true,
    algolia: {
      apiKey: '050bcd29a00cc5551793c2d6eadfb2a1',
      indexName: 'terraform-validator',
      algoliaOptions: {}, // Optional, if provided by Algolia
    },
    navbar: {
      title: '',
      logo: {
        alt: 'Terraform-Validator',
        src: 'img/terraform-validator.svg',
      },
      links: [
        {to: 'docs/getting-started/introduction', label: 'Docs', position: 'left'},
        {
          href: 'https://hub.docker.com/r/thazelart/terraform-validator',
          label: 'DockerHub',
          position: 'right',
        },
        {
          href: 'https://github.com/thazelart/terraform-validator',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [],
      copyright: `Copyright © ${new Date().getFullYear()} thazelart, built with Docusaurus.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
