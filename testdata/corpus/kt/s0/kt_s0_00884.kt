package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7250")
fun step7250(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
sealed class Outcome7251 {
    data class Success(val payload: String) : Outcome7251()
    data object Failure : Outcome7251()
}

fun describe7251(outcome: Outcome7251): String = when (outcome) {
    is Outcome7251.Success -> outcome.payload
    Outcome7251.Failure -> "failed"
}
sealed class Outcome7252 {
    data class Success(val payload: String) : Outcome7252()
    data object Failure : Outcome7252()
}

fun describe7252(outcome: Outcome7252): String = when (outcome) {
    is Outcome7252.Success -> outcome.payload
    Outcome7252.Failure -> "failed"
}
class Counter7253 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7254(): T? = if (size > 1) this[1] else null

fun render7254(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome7255 {
    data class Success(val payload: String) : Outcome7255()
    data object Failure : Outcome7255()
}

fun describe7255(outcome: Outcome7255): String = when (outcome) {
    is Outcome7255.Success -> outcome.payload
    Outcome7255.Failure -> "failed"
}
class Counter7256 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
