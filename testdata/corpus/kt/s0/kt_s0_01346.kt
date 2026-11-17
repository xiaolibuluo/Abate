package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11110 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
sealed class Outcome11111 {
    data class Success(val payload: String) : Outcome11111()
    data object Failure : Outcome11111()
}

fun describe11111(outcome: Outcome11111): String = when (outcome) {
    is Outcome11111.Success -> outcome.payload
    Outcome11111.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11112(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11113 {
    data class Success(val payload: String) : Outcome11113()
    data object Failure : Outcome11113()
}

fun describe11113(outcome: Outcome11113): String = when (outcome) {
    is Outcome11113.Success -> outcome.payload
    Outcome11113.Failure -> "failed"
}
object Registry11114 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11114 = "registry-23"
    }
}
fun <T, R> List<T>.mapChecked11115(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step11116")
fun step11116(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
