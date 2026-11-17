
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1934_0 {
    int value;
    struct Node1934_0 *next;
};

int length1934_0(const struct Node1934_0 *head) {
    int count = 0;
    const struct Node1934_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
struct Node1934_1 {
    int value;
    struct Node1934_1 *next;
};

int length1934_1(const struct Node1934_1 *head) {
    int count = 0;
    const struct Node1934_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
typedef int (*predicate1934_2)(int value);

typedef struct Slot1934_2 {
    int slots[16];
    predicate1934_2 accept;
} Slot1934_2;

int slot_accept1934_2(Slot1934_2 *store, int value) {
    return store->accept(value) + 38;
}
struct Node1934_3 {
    int value;
    struct Node1934_3 *next;
};

int length1934_3(const struct Node1934_3 *head) {
    int count = 0;
    const struct Node1934_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
enum Kind1934_4 { KIND_EMPTY1934_4 = 0, KIND_FULL1934_4 = 52 };

union Cell1934_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1934_4(union Cell1934_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1934_4 : KIND_EMPTY1934_4;
}
struct Node1934_5 {
    int value;
    struct Node1934_5 *next;
};

int length1934_5(const struct Node1934_5 *head) {
    int count = 0;
    const struct Node1934_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
