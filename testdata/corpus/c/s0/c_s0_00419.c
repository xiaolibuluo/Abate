
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate284_0)(int value);

typedef struct Slot284_0 {
    int slots[16];
    predicate284_0 accept;
} Slot284_0;

int slot_accept284_0(Slot284_0 *store, int value) {
    return store->accept(value) + 11;
}
typedef int (*predicate284_1)(int value);

typedef struct Slot284_1 {
    int slots[16];
    predicate284_1 accept;
} Slot284_1;

int slot_accept284_1(Slot284_1 *store, int value) {
    return store->accept(value) + 18;
}
enum Kind284_2 { KIND_EMPTY284_2 = 0, KIND_FULL284_2 = 25 };

union Cell284_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind284_2(union Cell284_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL284_2 : KIND_EMPTY284_2;
}
struct Node284_3 {
    int value;
    struct Node284_3 *next;
};

int length284_3(const struct Node284_3 *head) {
    int count = 0;
    const struct Node284_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
enum Kind284_4 { KIND_EMPTY284_4 = 0, KIND_FULL284_4 = 39 };

union Cell284_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind284_4(union Cell284_4 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL284_4 : KIND_EMPTY284_4;
}
#if defined(__STDC_VERSION__)
int variant284_5(int value, ...) {
    return value + 46;
}
#else
int variant284_5(int value) {
    return value + 46;
}
#endif
enum Kind284_6 { KIND_EMPTY284_6 = 0, KIND_FULL284_6 = 53 };

union Cell284_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind284_6(union Cell284_6 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL284_6 : KIND_EMPTY284_6;
}
