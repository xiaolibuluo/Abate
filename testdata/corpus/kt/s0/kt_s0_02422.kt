package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19910")
fun step19910(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
sealed class Outcome19911 {
    data class Success(val payload: String) : Outcome19911()
    data object Failure : Outcome19911()
}

fun describe19911(outcome: Outcome19911): String = when (outcome) {
    is Outcome19911.Success -> outcome.payload
    Outcome19911.Failure -> "failed"
}
class Counter19912 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
class Counter19913 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
@JvmName("step19914")
fun step19914(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
class Counter19915 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
object Registry19916 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19916 = "registry-31"
    }
}
