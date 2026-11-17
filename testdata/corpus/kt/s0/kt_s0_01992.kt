package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16470<A, B>(val first: A, val second: B)

fun combine16470(left: Int, right: Int): Pair16470<Int, Int> {
    return Pair16470(left + 76, right + 9)
}
object Registry16471 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16471 = "registry-83"
    }
}
object Registry16472 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16472 = "registry-90"
    }
}
object Registry16473 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16473 = "registry-97"
    }
}
class Counter16474 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
class Counter16475 {
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
fun <T> List<T>.secondOrNull16476(): T? = if (size > 1) this[1] else null

fun render16476(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16477 {
    data class Success(val payload: String) : Outcome16477()
    data object Failure : Outcome16477()
}

fun describe16477(outcome: Outcome16477): String = when (outcome) {
    is Outcome16477.Success -> outcome.payload
    Outcome16477.Failure -> "failed"
}
