
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1979_0 {
    int value;
    struct Node1979_0 *next;
};

int length1979_0(const struct Node1979_0 *head) {
    int count = 0;
    const struct Node1979_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
enum Kind1979_1 { KIND_EMPTY1979_1 = 0, KIND_FULL1979_1 = 34 };

union Cell1979_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1979_1(union Cell1979_1 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1979_1 : KIND_EMPTY1979_1;
}
struct Node1979_2 {
    int value;
    struct Node1979_2 *next;
};

int length1979_2(const struct Node1979_2 *head) {
    int count = 0;
    const struct Node1979_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
typedef int (*predicate1979_3)(int value);

typedef struct Slot1979_3 {
    int slots[16];
    predicate1979_3 accept;
} Slot1979_3;

int slot_accept1979_3(Slot1979_3 *store, int value) {
    return store->accept(value) + 48;
}
enum Kind1979_4 { KIND_EMPTY1979_4 = 0, KIND_FULL1979_4 = 55 };

union Cell1979_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1979_4(union Cell1979_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1979_4 : KIND_EMPTY1979_4;
}
