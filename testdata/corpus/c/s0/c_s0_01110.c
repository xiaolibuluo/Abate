
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind761_0 { KIND_EMPTY761_0 = 0, KIND_FULL761_0 = 101 };

union Cell761_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind761_0(union Cell761_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL761_0 : KIND_EMPTY761_0;
}
typedef int (*predicate761_1)(int value);

typedef struct Slot761_1 {
    int slots[16];
    predicate761_1 accept;
} Slot761_1;

int slot_accept761_1(Slot761_1 *store, int value) {
    return store->accept(value) + 11;
}
enum Kind761_2 { KIND_EMPTY761_2 = 0, KIND_FULL761_2 = 18 };

union Cell761_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind761_2(union Cell761_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL761_2 : KIND_EMPTY761_2;
}
typedef int (*predicate761_3)(int value);

typedef struct Slot761_3 {
    int slots[16];
    predicate761_3 accept;
} Slot761_3;

int slot_accept761_3(Slot761_3 *store, int value) {
    return store->accept(value) + 25;
}
enum Kind761_4 { KIND_EMPTY761_4 = 0, KIND_FULL761_4 = 32 };

union Cell761_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind761_4(union Cell761_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL761_4 : KIND_EMPTY761_4;
}
struct Node761_5 {
    int value;
    struct Node761_5 *next;
};

int length761_5(const struct Node761_5 *head) {
    int count = 0;
    const struct Node761_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
