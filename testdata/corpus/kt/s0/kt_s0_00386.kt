package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3180 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
@JvmName("step3181")
fun step3181(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
sealed class Outcome3182 {
    data class Success(val payload: String) : Outcome3182()
    data object Failure : Outcome3182()
}

fun describe3182(outcome: Outcome3182): String = when (outcome) {
    is Outcome3182.Success -> outcome.payload
    Outcome3182.Failure -> "failed"
}
sealed class Outcome3183 {
    data class Success(val payload: String) : Outcome3183()
    data object Failure : Outcome3183()
}

fun describe3183(outcome: Outcome3183): String = when (outcome) {
    is Outcome3183.Success -> outcome.payload
    Outcome3183.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3184(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
