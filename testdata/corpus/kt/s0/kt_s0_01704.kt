package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14020")
fun step14020(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
@JvmName("step14021")
fun step14021(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
data class Pair14022<A, B>(val first: A, val second: B)

fun combine14022(left: Int, right: Int): Pair14022<Int, Int> {
    return Pair14022(left + 9, right + 24)
}
class Counter14023 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14024(): T? = if (size > 1) this[1] else null

fun render14024(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14025 {
    data class Success(val payload: String) : Outcome14025()
    data object Failure : Outcome14025()
}

fun describe14025(outcome: Outcome14025): String = when (outcome) {
    is Outcome14025.Success -> outcome.payload
    Outcome14025.Failure -> "failed"
}
class Counter14026 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
