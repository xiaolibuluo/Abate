package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3390(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3391")
fun step3391(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
class Counter3392 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step3393")
fun step3393(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
fun <T> List<T>.secondOrNull3394(): T? = if (size > 1) this[1] else null

fun render3394(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3395 {
    data class Success(val payload: String) : Outcome3395()
    data object Failure : Outcome3395()
}

fun describe3395(outcome: Outcome3395): String = when (outcome) {
    is Outcome3395.Success -> outcome.payload
    Outcome3395.Failure -> "failed"
}
@JvmName("step3396")
fun step3396(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
data class Pair3397<A, B>(val first: A, val second: B)

fun combine3397(left: Int, right: Int): Pair3397<Int, Int> {
    return Pair3397(left + 96, right + 19)
}
sealed class Outcome3398 {
    data class Success(val payload: String) : Outcome3398()
    data object Failure : Outcome3398()
}

fun describe3398(outcome: Outcome3398): String = when (outcome) {
    is Outcome3398.Success -> outcome.payload
    Outcome3398.Failure -> "failed"
}
