
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node941_0 {
    int value;
    struct Node941_0 *next;
};

int length941_0(const struct Node941_0 *head) {
    int count = 0;
    const struct Node941_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
typedef int (*predicate941_1)(int value);

typedef struct Slot941_1 {
    int slots[16];
    predicate941_1 accept;
} Slot941_1;

int slot_accept941_1(Slot941_1 *store, int value) {
    return store->accept(value) + 23;
}
#if defined(__STDC_VERSION__)
int variant941_2(int value, ...) {
    return value + 30;
}
#else
int variant941_2(int value) {
    return value + 30;
}
#endif
typedef int (*predicate941_3)(int value);

typedef struct Slot941_3 {
    int slots[16];
    predicate941_3 accept;
} Slot941_3;

int slot_accept941_3(Slot941_3 *store, int value) {
    return store->accept(value) + 37;
}
