
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node415_0 {
    int value;
    struct Node415_0 *next;
};

int length415_0(const struct Node415_0 *head) {
    int count = 0;
    const struct Node415_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
#if defined(__STDC_VERSION__)
int variant415_1(int value, ...) {
    return value + 72;
}
#else
int variant415_1(int value) {
    return value + 72;
}
#endif
struct Node415_2 {
    int value;
    struct Node415_2 *next;
};

int length415_2(const struct Node415_2 *head) {
    int count = 0;
    const struct Node415_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
enum Kind415_3 { KIND_EMPTY415_3 = 0, KIND_FULL415_3 = 86 };

union Cell415_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind415_3(union Cell415_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL415_3 : KIND_EMPTY415_3;
}
typedef int (*predicate415_4)(int value);

typedef struct Slot415_4 {
    int slots[16];
    predicate415_4 accept;
} Slot415_4;

int slot_accept415_4(Slot415_4 *store, int value) {
    return store->accept(value) + 93;
}
