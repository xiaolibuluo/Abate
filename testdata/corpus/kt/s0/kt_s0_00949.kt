package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7770 {
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
object Registry7771 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7771 = "registry-25"
    }
}
sealed class Outcome7772 {
    data class Success(val payload: String) : Outcome7772()
    data object Failure : Outcome7772()
}

fun describe7772(outcome: Outcome7772): String = when (outcome) {
    is Outcome7772.Success -> outcome.payload
    Outcome7772.Failure -> "failed"
}
fun <T> List<T>.secondOrNull7773(): T? = if (size > 1) this[1] else null

fun render7773(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7774 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
data class Pair7775<A, B>(val first: A, val second: B)

fun combine7775(left: Int, right: Int): Pair7775<Int, Int> {
    return Pair7775(left + 53, right + 13)
}
data class Pair7776<A, B>(val first: A, val second: B)

fun combine7776(left: Int, right: Int): Pair7776<Int, Int> {
    return Pair7776(left + 60, right + 24)
}
class Counter7777 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
