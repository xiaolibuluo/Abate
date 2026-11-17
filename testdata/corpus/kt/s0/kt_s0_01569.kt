package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12990 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
@JvmName("step12991")
fun step12991(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
object Registry12992 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12992 = "registry-28"
    }
}
sealed class Outcome12993 {
    data class Success(val payload: String) : Outcome12993()
    data object Failure : Outcome12993()
}

fun describe12993(outcome: Outcome12993): String = when (outcome) {
    is Outcome12993.Success -> outcome.payload
    Outcome12993.Failure -> "failed"
}
class Counter12994 {
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
object Registry12995 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12995 = "registry-49"
    }
}
