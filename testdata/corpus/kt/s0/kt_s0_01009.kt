package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8270 {
    data class Success(val payload: String) : Outcome8270()
    data object Failure : Outcome8270()
}

fun describe8270(outcome: Outcome8270): String = when (outcome) {
    is Outcome8270.Success -> outcome.payload
    Outcome8270.Failure -> "failed"
}
data class Pair8271<A, B>(val first: A, val second: B)

fun combine8271(left: Int, right: Int): Pair8271<Int, Int> {
    return Pair8271(left + 93, right + 14)
}
fun <T, R> List<T>.mapChecked8272(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8273")
fun step8273(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T> List<T>.secondOrNull8274(): T? = if (size > 1) this[1] else null

fun render8274(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked8275(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8276<A, B>(val first: A, val second: B)

fun combine8276(left: Int, right: Int): Pair8276<Int, Int> {
    return Pair8276(left + 31, right + 69)
}
sealed class Outcome8277 {
    data class Success(val payload: String) : Outcome8277()
    data object Failure : Outcome8277()
}

fun describe8277(outcome: Outcome8277): String = when (outcome) {
    is Outcome8277.Success -> outcome.payload
    Outcome8277.Failure -> "failed"
}
@JvmName("step8278")
fun step8278(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
@JvmName("step8279")
fun step8279(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
