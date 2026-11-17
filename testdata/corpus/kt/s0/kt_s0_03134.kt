package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25720 {
    data class Success(val payload: String) : Outcome25720()
    data object Failure : Outcome25720()
}

fun describe25720(outcome: Outcome25720): String = when (outcome) {
    is Outcome25720.Success -> outcome.payload
    Outcome25720.Failure -> "failed"
}
sealed class Outcome25721 {
    data class Success(val payload: String) : Outcome25721()
    data object Failure : Outcome25721()
}

fun describe25721(outcome: Outcome25721): String = when (outcome) {
    is Outcome25721.Success -> outcome.payload
    Outcome25721.Failure -> "failed"
}
class Counter25722 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25723(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step25724")
fun step25724(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
@JvmName("step25725")
fun step25725(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
class Counter25726 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
