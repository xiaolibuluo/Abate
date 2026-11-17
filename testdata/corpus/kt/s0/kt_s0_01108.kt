package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked9080(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome9081 {
    data class Success(val payload: String) : Outcome9081()
    data object Failure : Outcome9081()
}

fun describe9081(outcome: Outcome9081): String = when (outcome) {
    is Outcome9081.Success -> outcome.payload
    Outcome9081.Failure -> "failed"
}
fun <T> List<T>.secondOrNull9082(): T? = if (size > 1) this[1] else null

fun render9082(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked9083(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9084<A, B>(val first: A, val second: B)

fun combine9084(left: Int, right: Int): Pair9084<Int, Int> {
    return Pair9084(left + 100, right + 65)
}
@JvmName("step9085")
fun step9085(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
sealed class Outcome9086 {
    data class Success(val payload: String) : Outcome9086()
    data object Failure : Outcome9086()
}

fun describe9086(outcome: Outcome9086): String = when (outcome) {
    is Outcome9086.Success -> outcome.payload
    Outcome9086.Failure -> "failed"
}
