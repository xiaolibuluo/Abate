package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9210")
fun step9210(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
data class Pair9211<A, B>(val first: A, val second: B)

fun combine9211(left: Int, right: Int): Pair9211<Int, Int> {
    return Pair9211(left + 54, right + 62)
}
sealed class Outcome9212 {
    data class Success(val payload: String) : Outcome9212()
    data object Failure : Outcome9212()
}

fun describe9212(outcome: Outcome9212): String = when (outcome) {
    is Outcome9212.Success -> outcome.payload
    Outcome9212.Failure -> "failed"
}
data class Pair9213<A, B>(val first: A, val second: B)

fun combine9213(left: Int, right: Int): Pair9213<Int, Int> {
    return Pair9213(left + 68, right + 23)
}
@JvmName("step9214")
fun step9214(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
fun <T> List<T>.secondOrNull9215(): T? = if (size > 1) this[1] else null

fun render9215(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9216")
fun step9216(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
fun <T, R> List<T>.mapChecked9217(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9218 {
    data class Success(val payload: String) : Outcome9218()
    data object Failure : Outcome9218()
}

fun describe9218(outcome: Outcome9218): String = when (outcome) {
    is Outcome9218.Success -> outcome.payload
    Outcome9218.Failure -> "failed"
}
sealed class Outcome9219 {
    data class Success(val payload: String) : Outcome9219()
    data object Failure : Outcome9219()
}

fun describe9219(outcome: Outcome9219): String = when (outcome) {
    is Outcome9219.Success -> outcome.payload
    Outcome9219.Failure -> "failed"
}
