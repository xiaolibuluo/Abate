
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node642_0 {
    int value;
    struct Node642_0 *next;
};

int length642_0(const struct Node642_0 *head) {
    int count = 0;
    const struct Node642_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
enum Kind642_1 { KIND_EMPTY642_1 = 0, KIND_FULL642_1 = 16 };

union Cell642_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind642_1(union Cell642_1 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL642_1 : KIND_EMPTY642_1;
}
