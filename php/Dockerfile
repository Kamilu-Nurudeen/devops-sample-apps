FROM php:8.2-apache

WORKDIR /var/www/html

# Copy application files
COPY index.php ./
COPY config.prod ./config

ENV APP_ENV=prod

EXPOSE 80

CMD ["apache2-foreground"] 
