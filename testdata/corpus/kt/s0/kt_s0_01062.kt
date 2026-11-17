package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8730 = "registry-5"
    }
}
object Registry8731 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8731 = "registry-12"
    }
}
object Registry8732 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8732 = "registry-19"
    }
}
sealed class Outcome8733 {
    data class Success(val payload: String) : Outcome8733()
    data object Failure : Outcome8733()
}

fun describe8733(outcome: Outcome8733): String = when (outcome) {
    is Outcome8733.Success -> outcome.payload
    Outcome8733.Failure -> "failed"
}
data class Pair8734<A, B>(val first: A, val second: B)

fun combine8734(left: Int, right: Int): Pair8734<Int, Int> {
    return Pair8734(left + 33, right + 64)
}
fun <T> List<T>.secondOrNull8735(): T? = if (size > 1) this[1] else null

fun render8735(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8736 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8736 = "registry-47"
    }
}
@JvmName("step8737")
fun step8737(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
fun <T> List<T>.secondOrNull8738(): T? = if (size > 1) this[1] else null

fun render8738(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8739 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
