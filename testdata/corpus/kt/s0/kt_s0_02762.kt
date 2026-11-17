package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22810 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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
sealed class Outcome22811 {
    data class Success(val payload: String) : Outcome22811()
    data object Failure : Outcome22811()
}

fun describe22811(outcome: Outcome22811): String = when (outcome) {
    is Outcome22811.Success -> outcome.payload
    Outcome22811.Failure -> "failed"
}
@JvmName("step22812")
fun step22812(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
fun <T> List<T>.secondOrNull22813(): T? = if (size > 1) this[1] else null

fun render22813(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22814 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22814 = "registry-101"
    }
}
class Counter22815 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
class Counter22816 {
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
