package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter4700 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
@JvmName("step4701")
fun step4701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
sealed class Outcome4702 {
    data class Success(val payload: String) : Outcome4702()
    data object Failure : Outcome4702()
}

fun describe4702(outcome: Outcome4702): String = when (outcome) {
    is Outcome4702.Success -> outcome.payload
    Outcome4702.Failure -> "failed"
}
@JvmName("step4703")
fun step4703(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
@JvmName("step4704")
fun step4704(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
fun <T> List<T>.secondOrNull4705(): T? = if (size > 1) this[1] else null

fun render4705(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
