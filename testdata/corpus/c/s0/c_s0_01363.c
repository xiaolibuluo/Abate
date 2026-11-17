
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node936_0 {
    int value;
    struct Node936_0 *next;
};

int length936_0(const struct Node936_0 *head) {
    int count = 0;
    const struct Node936_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
enum Kind936_1 { KIND_EMPTY936_1 = 0, KIND_FULL936_1 = 55 };

union Cell936_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind936_1(union Cell936_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL936_1 : KIND_EMPTY936_1;
}
