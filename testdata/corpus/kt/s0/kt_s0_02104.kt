package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter17370 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
object Registry17371 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17371 = "registry-89"
    }
}
sealed class Outcome17372 {
    data class Success(val payload: String) : Outcome17372()
    data object Failure : Outcome17372()
}

fun describe17372(outcome: Outcome17372): String = when (outcome) {
    is Outcome17372.Success -> outcome.payload
    Outcome17372.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17373(): T? = if (size > 1) this[1] else null

fun render17373(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17374 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
