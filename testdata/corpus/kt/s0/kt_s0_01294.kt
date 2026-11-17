package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10680 {
    data class Success(val payload: String) : Outcome10680()
    data object Failure : Outcome10680()
}

fun describe10680(outcome: Outcome10680): String = when (outcome) {
    is Outcome10680.Success -> outcome.payload
    Outcome10680.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10681(): T? = if (size > 1) this[1] else null

fun render10681(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked10682(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry10683 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10683 = "registry-39"
    }
}
object Registry10684 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10684 = "registry-46"
    }
}
class Counter10685 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
data class Pair10686<A, B>(val first: A, val second: B)

fun combine10686(left: Int, right: Int): Pair10686<Int, Int> {
    return Pair10686(left + 60, right + 48)
}
object Registry10687 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10687 = "registry-67"
    }
}
class Counter10688 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
