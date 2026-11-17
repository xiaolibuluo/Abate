
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2509_0)(int value);

typedef struct Slot2509_0 {
    int slots[16];
    predicate2509_0 accept;
} Slot2509_0;

int slot_accept2509_0(Slot2509_0 *store, int value) {
    return store->accept(value) + 30;
}
typedef int (*predicate2509_1)(int value);

typedef struct Slot2509_1 {
    int slots[16];
    predicate2509_1 accept;
} Slot2509_1;

int slot_accept2509_1(Slot2509_1 *store, int value) {
    return store->accept(value) + 37;
}
#if defined(__STDC_VERSION__)
int variant2509_2(int value, ...) {
    return value + 44;
}
#else
int variant2509_2(int value) {
    return value + 44;
}
#endif
struct Node2509_3 {
    int value;
    struct Node2509_3 *next;
};

int length2509_3(const struct Node2509_3 *head) {
    int count = 0;
    const struct Node2509_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
#if defined(__STDC_VERSION__)
int variant2509_4(int value, ...) {
    return value + 58;
}
#else
int variant2509_4(int value) {
    return value + 58;
}
#endif
struct Node2509_5 {
    int value;
    struct Node2509_5 *next;
};

int length2509_5(const struct Node2509_5 *head) {
    int count = 0;
    const struct Node2509_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
