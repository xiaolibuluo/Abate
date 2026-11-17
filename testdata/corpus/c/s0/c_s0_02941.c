
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1988_0 { KIND_EMPTY1988_0 = 0, KIND_FULL1988_0 = 47 };

union Cell1988_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1988_0(union Cell1988_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1988_0 : KIND_EMPTY1988_0;
}
typedef int (*predicate1988_1)(int value);

typedef struct Slot1988_1 {
    int slots[16];
    predicate1988_1 accept;
} Slot1988_1;

int slot_accept1988_1(Slot1988_1 *store, int value) {
    return store->accept(value) + 54;
}
enum Kind1988_2 { KIND_EMPTY1988_2 = 0, KIND_FULL1988_2 = 61 };

union Cell1988_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1988_2(union Cell1988_2 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1988_2 : KIND_EMPTY1988_2;
}
struct Node1988_3 {
    int value;
    struct Node1988_3 *next;
};

int length1988_3(const struct Node1988_3 *head) {
    int count = 0;
    const struct Node1988_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
struct Node1988_4 {
    int value;
    struct Node1988_4 *next;
};

int length1988_4(const struct Node1988_4 *head) {
    int count = 0;
    const struct Node1988_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant1988_5(int value, ...) {
    return value + 82;
}
#else
int variant1988_5(int value) {
    return value + 82;
}
#endif
typedef int (*predicate1988_6)(int value);

typedef struct Slot1988_6 {
    int slots[16];
    predicate1988_6 accept;
} Slot1988_6;

int slot_accept1988_6(Slot1988_6 *store, int value) {
    return store->accept(value) + 89;
}
