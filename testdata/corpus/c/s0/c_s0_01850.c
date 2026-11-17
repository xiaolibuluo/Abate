
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1241_0 {
    int value;
    struct Node1241_0 *next;
};

int length1241_0(const struct Node1241_0 *head) {
    int count = 0;
    const struct Node1241_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
enum Kind1241_1 { KIND_EMPTY1241_1 = 0, KIND_FULL1241_1 = 43 };

union Cell1241_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1241_1(union Cell1241_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1241_1 : KIND_EMPTY1241_1;
}
enum Kind1241_2 { KIND_EMPTY1241_2 = 0, KIND_FULL1241_2 = 50 };

union Cell1241_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1241_2(union Cell1241_2 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1241_2 : KIND_EMPTY1241_2;
}
typedef int (*predicate1241_3)(int value);

typedef struct Slot1241_3 {
    int slots[16];
    predicate1241_3 accept;
} Slot1241_3;

int slot_accept1241_3(Slot1241_3 *store, int value) {
    return store->accept(value) + 57;
}
