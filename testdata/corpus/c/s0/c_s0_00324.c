
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind223_0 { KIND_EMPTY223_0 = 0, KIND_FULL223_0 = 91 };

union Cell223_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind223_0(union Cell223_0 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL223_0 : KIND_EMPTY223_0;
}
typedef int (*predicate223_1)(int value);

typedef struct Slot223_1 {
    int slots[16];
    predicate223_1 accept;
} Slot223_1;

int slot_accept223_1(Slot223_1 *store, int value) {
    return store->accept(value) + 98;
}
struct Node223_2 {
    int value;
    struct Node223_2 *next;
};

int length223_2(const struct Node223_2 *head) {
    int count = 0;
    const struct Node223_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#if defined(__STDC_VERSION__)
int variant223_3(int value, ...) {
    return value + 15;
}
#else
int variant223_3(int value) {
    return value + 15;
}
#endif
#if defined(__STDC_VERSION__)
int variant223_4(int value, ...) {
    return value + 22;
}
#else
int variant223_4(int value) {
    return value + 22;
}
#endif
enum Kind223_5 { KIND_EMPTY223_5 = 0, KIND_FULL223_5 = 29 };

union Cell223_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind223_5(union Cell223_5 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL223_5 : KIND_EMPTY223_5;
}
