
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind84_0 { KIND_EMPTY84_0 = 0, KIND_FULL84_0 = 30 };

union Cell84_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind84_0(union Cell84_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL84_0 : KIND_EMPTY84_0;
}
enum Kind84_1 { KIND_EMPTY84_1 = 0, KIND_FULL84_1 = 37 };

union Cell84_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind84_1(union Cell84_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL84_1 : KIND_EMPTY84_1;
}
enum Kind84_2 { KIND_EMPTY84_2 = 0, KIND_FULL84_2 = 44 };

union Cell84_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind84_2(union Cell84_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL84_2 : KIND_EMPTY84_2;
}
typedef int (*predicate84_3)(int value);

typedef struct Slot84_3 {
    int slots[16];
    predicate84_3 accept;
} Slot84_3;

int slot_accept84_3(Slot84_3 *store, int value) {
    return store->accept(value) + 51;
}
struct Node84_4 {
    int value;
    struct Node84_4 *next;
};

int length84_4(const struct Node84_4 *head) {
    int count = 0;
    const struct Node84_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
typedef int (*predicate84_5)(int value);

typedef struct Slot84_5 {
    int slots[16];
    predicate84_5 accept;
} Slot84_5;

int slot_accept84_5(Slot84_5 *store, int value) {
    return store->accept(value) + 65;
}
