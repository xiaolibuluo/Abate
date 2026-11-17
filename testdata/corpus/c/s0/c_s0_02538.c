
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1726_0)(int value);

typedef struct Slot1726_0 {
    int slots[16];
    predicate1726_0 accept;
} Slot1726_0;

int slot_accept1726_0(Slot1726_0 *store, int value) {
    return store->accept(value) + 36;
}
struct Node1726_1 {
    int value;
    struct Node1726_1 *next;
};

int length1726_1(const struct Node1726_1 *head) {
    int count = 0;
    const struct Node1726_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
enum Kind1726_2 { KIND_EMPTY1726_2 = 0, KIND_FULL1726_2 = 50 };

union Cell1726_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1726_2(union Cell1726_2 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1726_2 : KIND_EMPTY1726_2;
}
typedef int (*predicate1726_3)(int value);

typedef struct Slot1726_3 {
    int slots[16];
    predicate1726_3 accept;
} Slot1726_3;

int slot_accept1726_3(Slot1726_3 *store, int value) {
    return store->accept(value) + 57;
}
