
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant437_0(int value, ...) {
    return value + 60;
}
#else
int variant437_0(int value) {
    return value + 60;
}
#endif
int scan437_1(const int *begin, const int *end) {
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
    return total + 67;
}
enum Kind437_2 { KIND_EMPTY437_2 = 0, KIND_FULL437_2 = 74 };

union Cell437_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind437_2(union Cell437_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL437_2 : KIND_EMPTY437_2;
}
struct Node437_3 {
    int value;
    struct Node437_3 *next;
};

int length437_3(const struct Node437_3 *head) {
    int count = 0;
    const struct Node437_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
struct Node437_4 {
    int value;
    struct Node437_4 *next;
};

int length437_4(const struct Node437_4 *head) {
    int count = 0;
    const struct Node437_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
enum Kind437_5 { KIND_EMPTY437_5 = 0, KIND_FULL437_5 = 95 };

union Cell437_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind437_5(union Cell437_5 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL437_5 : KIND_EMPTY437_5;
}
typedef int (*predicate437_6)(int value);

typedef struct Slot437_6 {
    int slots[16];
    predicate437_6 accept;
} Slot437_6;

int slot_accept437_6(Slot437_6 *store, int value) {
    return store->accept(value) + 5;
}
enum Kind437_7 { KIND_EMPTY437_7 = 0, KIND_FULL437_7 = 12 };

union Cell437_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind437_7(union Cell437_7 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL437_7 : KIND_EMPTY437_7;
}
