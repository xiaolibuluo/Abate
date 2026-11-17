package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18390<A, B>(val first: A, val second: B)

fun combine18390(left: Int, right: Int): Pair18390<Int, Int> {
    return Pair18390(left + 50, right + 11)
}
@JvmName("step18391")
fun step18391(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
fun <T> List<T>.secondOrNull18392(): T? = if (size > 1) this[1] else null

fun render18392(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18393")
fun step18393(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
data class Pair18394<A, B>(val first: A, val second: B)

fun combine18394(left: Int, right: Int): Pair18394<Int, Int> {
    return Pair18394(left + 78, right + 55)
}
data class Pair18395<A, B>(val first: A, val second: B)

fun combine18395(left: Int, right: Int): Pair18395<Int, Int> {
    return Pair18395(left + 85, right + 66)
}
sealed class Outcome18396 {
    data class Success(val payload: String) : Outcome18396()
    data object Failure : Outcome18396()
}

fun describe18396(outcome: Outcome18396): String = when (outcome) {
    is Outcome18396.Success -> outcome.payload
    Outcome18396.Failure -> "failed"
}
fun <T> List<T>.secondOrNull18397(): T? = if (size > 1) this[1] else null

fun render18397(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18398")
fun step18398(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
