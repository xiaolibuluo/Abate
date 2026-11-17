package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8100<A, B>(val first: A, val second: B)

fun combine8100(left: Int, right: Int): Pair8100<Int, Int> {
    return Pair8100(left + 59, right + 67)
}
data class Pair8101<A, B>(val first: A, val second: B)

fun combine8101(left: Int, right: Int): Pair8101<Int, Int> {
    return Pair8101(left + 66, right + 17)
}
fun <T> List<T>.secondOrNull8102(): T? = if (size > 1) this[1] else null

fun render8102(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8103<A, B>(val first: A, val second: B)

fun combine8103(left: Int, right: Int): Pair8103<Int, Int> {
    return Pair8103(left + 80, right + 39)
}
data class Pair8104<A, B>(val first: A, val second: B)

fun combine8104(left: Int, right: Int): Pair8104<Int, Int> {
    return Pair8104(left + 87, right + 50)
}
sealed class Outcome8105 {
    data class Success(val payload: String) : Outcome8105()
    data object Failure : Outcome8105()
}

fun describe8105(outcome: Outcome8105): String = when (outcome) {
    is Outcome8105.Success -> outcome.payload
    Outcome8105.Failure -> "failed"
}
data class Pair8106<A, B>(val first: A, val second: B)

fun combine8106(left: Int, right: Int): Pair8106<Int, Int> {
    return Pair8106(left + 101, right + 11)
}
