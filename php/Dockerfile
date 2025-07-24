FROM php:8.2-apache

WORKDIR /var/www/html

# Enable Apache rewrite module
RUN a2enmod rewrite

# Copy application files
COPY index.php ./
COPY config.prod ./config
COPY .htaccess ./


ENV APP_ENV=prod

EXPOSE 80

CMD ["apache2-foreground"] 
