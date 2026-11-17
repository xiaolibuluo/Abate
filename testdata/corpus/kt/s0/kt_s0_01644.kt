package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13580 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
class Counter13581 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
class Counter13582 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
data class Pair13583<A, B>(val first: A, val second: B)

fun combine13583(left: Int, right: Int): Pair13583<Int, Int> {
    return Pair13583(left + 26, right + 32)
}
sealed class Outcome13584 {
    data class Success(val payload: String) : Outcome13584()
    data object Failure : Outcome13584()
}

fun describe13584(outcome: Outcome13584): String = when (outcome) {
    is Outcome13584.Success -> outcome.payload
    Outcome13584.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13585(): T? = if (size > 1) this[1] else null

fun render13585(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13586 {
    data class Success(val payload: String) : Outcome13586()
    data object Failure : Outcome13586()
}

fun describe13586(outcome: Outcome13586): String = when (outcome) {
    is Outcome13586.Success -> outcome.payload
    Outcome13586.Failure -> "failed"
}
