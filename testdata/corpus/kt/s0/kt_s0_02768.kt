package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22850 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
object Registry22851 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22851 = "registry-35"
    }
}
class Counter22852 {
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
fun <T> List<T>.secondOrNull22853(): T? = if (size > 1) this[1] else null

fun render22853(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step22854")
fun step22854(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
data class Pair22855<A, B>(val first: A, val second: B)

fun combine22855(left: Int, right: Int): Pair22855<Int, Int> {
    return Pair22855(left + 63, right + 16)
}
sealed class Outcome22856 {
    data class Success(val payload: String) : Outcome22856()
    data object Failure : Outcome22856()
}

fun describe22856(outcome: Outcome22856): String = when (outcome) {
    is Outcome22856.Success -> outcome.payload
    Outcome22856.Failure -> "failed"
}
object Registry22857 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22857 = "registry-77"
    }
}
