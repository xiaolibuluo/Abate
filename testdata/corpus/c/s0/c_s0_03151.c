
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2135_0 {
    int value;
    struct Node2135_0 *next;
};

int length2135_0(const struct Node2135_0 *head) {
    int count = 0;
    const struct Node2135_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
typedef int (*predicate2135_1)(int value);

typedef struct Slot2135_1 {
    int slots[16];
    predicate2135_1 accept;
} Slot2135_1;

int slot_accept2135_1(Slot2135_1 *store, int value) {
    return store->accept(value) + 25;
}
struct Node2135_2 {
    int value;
    struct Node2135_2 *next;
};

int length2135_2(const struct Node2135_2 *head) {
    int count = 0;
    const struct Node2135_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
enum Kind2135_3 { KIND_EMPTY2135_3 = 0, KIND_FULL2135_3 = 39 };

union Cell2135_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2135_3(union Cell2135_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL2135_3 : KIND_EMPTY2135_3;
}
#if defined(__STDC_VERSION__)
int variant2135_4(int value, ...) {
    return value + 46;
}
#else
int variant2135_4(int value) {
    return value + 46;
}
#endif
int scan2135_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 53;
}
struct Node2135_6 {
    int value;
    struct Node2135_6 *next;
};

int length2135_6(const struct Node2135_6 *head) {
    int count = 0;
    const struct Node2135_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
