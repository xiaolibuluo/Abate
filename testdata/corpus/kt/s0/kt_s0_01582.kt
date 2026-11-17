package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13110<A, B>(val first: A, val second: B)

fun combine13110(left: Int, right: Int): Pair13110<Int, Int> {
    return Pair13110(left + 73, right + 36)
}
data class Pair13111<A, B>(val first: A, val second: B)

fun combine13111(left: Int, right: Int): Pair13111<Int, Int> {
    return Pair13111(left + 80, right + 47)
}
sealed class Outcome13112 {
    data class Success(val payload: String) : Outcome13112()
    data object Failure : Outcome13112()
}

fun describe13112(outcome: Outcome13112): String = when (outcome) {
    is Outcome13112.Success -> outcome.payload
    Outcome13112.Failure -> "failed"
}
data class Pair13113<A, B>(val first: A, val second: B)

fun combine13113(left: Int, right: Int): Pair13113<Int, Int> {
    return Pair13113(left + 94, right + 69)
}
data class Pair13114<A, B>(val first: A, val second: B)

fun combine13114(left: Int, right: Int): Pair13114<Int, Int> {
    return Pair13114(left + 101, right + 19)
}
fun <T, R> List<T>.mapChecked13115(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13116 {
    data class Success(val payload: String) : Outcome13116()
    data object Failure : Outcome13116()
}

fun describe13116(outcome: Outcome13116): String = when (outcome) {
    is Outcome13116.Success -> outcome.payload
    Outcome13116.Failure -> "failed"
}
data class Pair13117<A, B>(val first: A, val second: B)

fun combine13117(left: Int, right: Int): Pair13117<Int, Int> {
    return Pair13117(left + 25, right + 52)
}
fun <T> List<T>.secondOrNull13118(): T? = if (size > 1) this[1] else null

fun render13118(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
