
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant632_0(int value, ...) {
    return value + 73;
}
#else
int variant632_0(int value) {
    return value + 73;
}
#endif
enum Kind632_1 { KIND_EMPTY632_1 = 0, KIND_FULL632_1 = 80 };

union Cell632_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind632_1(union Cell632_1 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL632_1 : KIND_EMPTY632_1;
}
enum Kind632_2 { KIND_EMPTY632_2 = 0, KIND_FULL632_2 = 87 };

union Cell632_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind632_2(union Cell632_2 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL632_2 : KIND_EMPTY632_2;
}
typedef int (*predicate632_3)(int value);

typedef struct Slot632_3 {
    int slots[16];
    predicate632_3 accept;
} Slot632_3;

int slot_accept632_3(Slot632_3 *store, int value) {
    return store->accept(value) + 94;
}
struct Node632_4 {
    int value;
    struct Node632_4 *next;
};

int length632_4(const struct Node632_4 *head) {
    int count = 0;
    const struct Node632_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
enum Kind632_5 { KIND_EMPTY632_5 = 0, KIND_FULL632_5 = 11 };

union Cell632_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind632_5(union Cell632_5 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL632_5 : KIND_EMPTY632_5;
}
#if defined(__STDC_VERSION__)
int variant632_6(int value, ...) {
    return value + 18;
}
#else
int variant632_6(int value) {
    return value + 18;
}
#endif
struct Node632_7 {
    int value;
    struct Node632_7 *next;
};

int length632_7(const struct Node632_7 *head) {
    int count = 0;
    const struct Node632_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
