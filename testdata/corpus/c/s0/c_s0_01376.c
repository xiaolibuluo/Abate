
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node945_0 {
    int value;
    struct Node945_0 *next;
};

int length945_0(const struct Node945_0 *head) {
    int count = 0;
    const struct Node945_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#if defined(__STDC_VERSION__)
int variant945_1(int value, ...) {
    return value + 75;
}
#else
int variant945_1(int value) {
    return value + 75;
}
#endif
struct Node945_2 {
    int value;
    struct Node945_2 *next;
};

int length945_2(const struct Node945_2 *head) {
    int count = 0;
    const struct Node945_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
#if defined(__STDC_VERSION__)
int variant945_3(int value, ...) {
    return value + 89;
}
#else
int variant945_3(int value) {
    return value + 89;
}
#endif
