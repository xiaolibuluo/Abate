package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome930 {
    data class Success(val payload: String) : Outcome930()
    data object Failure : Outcome930()
}

fun describe930(outcome: Outcome930): String = when (outcome) {
    is Outcome930.Success -> outcome.payload
    Outcome930.Failure -> "failed"
}
data class Pair931<A, B>(val first: A, val second: B)

fun combine931(left: Int, right: Int): Pair931<Int, Int> {
    return Pair931(left + 57, right + 61)
}
class Counter932 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
object Registry933 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG933 = "registry-71"
    }
}
class Counter934 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull935(): T? = if (size > 1) this[1] else null

fun render935(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry936 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG936 = "registry-92"
    }
}
object Registry937 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG937 = "registry-99"
    }
}
