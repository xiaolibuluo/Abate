package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18960 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
object Registry18961 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18961 = "registry-22"
    }
}
sealed class Outcome18962 {
    data class Success(val payload: String) : Outcome18962()
    data object Failure : Outcome18962()
}

fun describe18962(outcome: Outcome18962): String = when (outcome) {
    is Outcome18962.Success -> outcome.payload
    Outcome18962.Failure -> "failed"
}
object Registry18963 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18963 = "registry-36"
    }
}
class Counter18964 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
data class Pair18965<A, B>(val first: A, val second: B)

fun combine18965(left: Int, right: Int): Pair18965<Int, Int> {
    return Pair18965(left + 50, right + 38)
}
fun <T> List<T>.secondOrNull18966(): T? = if (size > 1) this[1] else null

fun render18966(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter18967 {
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
