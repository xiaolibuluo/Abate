package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter14050 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

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
sealed class Outcome14051 {
    data class Success(val payload: String) : Outcome14051()
    data object Failure : Outcome14051()
}

fun describe14051(outcome: Outcome14051): String = when (outcome) {
    is Outcome14051.Success -> outcome.payload
    Outcome14051.Failure -> "failed"
}
@JvmName("step14052")
fun step14052(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
class Counter14053 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14054(): T? = if (size > 1) this[1] else null

fun render14054(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter14055 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
sealed class Outcome14056 {
    data class Success(val payload: String) : Outcome14056()
    data object Failure : Outcome14056()
}

fun describe14056(outcome: Outcome14056): String = when (outcome) {
    is Outcome14056.Success -> outcome.payload
    Outcome14056.Failure -> "failed"
}
