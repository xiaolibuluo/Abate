
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2501_0(int value, ...) {
    return value + 23;
}
#else
int variant2501_0(int value) {
    return value + 23;
}
#endif
#if defined(__STDC_VERSION__)
int variant2501_1(int value, ...) {
    return value + 30;
}
#else
int variant2501_1(int value) {
    return value + 30;
}
#endif
struct Node2501_2 {
    int value;
    struct Node2501_2 *next;
};

int length2501_2(const struct Node2501_2 *head) {
    int count = 0;
    const struct Node2501_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
struct Node2501_3 {
    int value;
    struct Node2501_3 *next;
};

int length2501_3(const struct Node2501_3 *head) {
    int count = 0;
    const struct Node2501_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
