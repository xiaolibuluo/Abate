package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12130 {
    data class Success(val payload: String) : Outcome12130()
    data object Failure : Outcome12130()
}

fun describe12130(outcome: Outcome12130): String = when (outcome) {
    is Outcome12130.Success -> outcome.payload
    Outcome12130.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12131(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter12132 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
sealed class Outcome12133 {
    data class Success(val payload: String) : Outcome12133()
    data object Failure : Outcome12133()
}

fun describe12133(outcome: Outcome12133): String = when (outcome) {
    is Outcome12133.Success -> outcome.payload
    Outcome12133.Failure -> "failed"
}
class Counter12134 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
class Counter12135 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
