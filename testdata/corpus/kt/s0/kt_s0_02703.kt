package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22330 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull22331(): T? = if (size > 1) this[1] else null

fun render22331(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22332 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
sealed class Outcome22333 {
    data class Success(val payload: String) : Outcome22333()
    data object Failure : Outcome22333()
}

fun describe22333(outcome: Outcome22333): String = when (outcome) {
    is Outcome22333.Success -> outcome.payload
    Outcome22333.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22334(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22335 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
sealed class Outcome22336 {
    data class Success(val payload: String) : Outcome22336()
    data object Failure : Outcome22336()
}

fun describe22336(outcome: Outcome22336): String = when (outcome) {
    is Outcome22336.Success -> outcome.payload
    Outcome22336.Failure -> "failed"
}
