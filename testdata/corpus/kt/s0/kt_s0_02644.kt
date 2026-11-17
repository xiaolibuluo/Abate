package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21850 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
@JvmName("step21851")
fun step21851(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
data class Pair21852<A, B>(val first: A, val second: B)

fun combine21852(left: Int, right: Int): Pair21852<Int, Int> {
    return Pair21852(left + 100, right + 15)
}
sealed class Outcome21853 {
    data class Success(val payload: String) : Outcome21853()
    data object Failure : Outcome21853()
}

fun describe21853(outcome: Outcome21853): String = when (outcome) {
    is Outcome21853.Success -> outcome.payload
    Outcome21853.Failure -> "failed"
}
data class Pair21854<A, B>(val first: A, val second: B)

fun combine21854(left: Int, right: Int): Pair21854<Int, Int> {
    return Pair21854(left + 17, right + 37)
}
object Registry21855 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21855 = "registry-24"
    }
}
data class Pair21856<A, B>(val first: A, val second: B)

fun combine21856(left: Int, right: Int): Pair21856<Int, Int> {
    return Pair21856(left + 31, right + 59)
}
fun <T> List<T>.secondOrNull21857(): T? = if (size > 1) this[1] else null

fun render21857(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome21858 {
    data class Success(val payload: String) : Outcome21858()
    data object Failure : Outcome21858()
}

fun describe21858(outcome: Outcome21858): String = when (outcome) {
    is Outcome21858.Success -> outcome.payload
    Outcome21858.Failure -> "failed"
}
