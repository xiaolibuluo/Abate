package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20890 {
    data class Success(val payload: String) : Outcome20890()
    data object Failure : Outcome20890()
}

fun describe20890(outcome: Outcome20890): String = when (outcome) {
    is Outcome20890.Success -> outcome.payload
    Outcome20890.Failure -> "failed"
}
data class Pair20891<A, B>(val first: A, val second: B)

fun combine20891(left: Int, right: Int): Pair20891<Int, Int> {
    return Pair20891(left + 9, right + 64)
}
object Registry20892 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20892 = "registry-16"
    }
}
fun <T> List<T>.secondOrNull20893(): T? = if (size > 1) this[1] else null

fun render20893(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter20894 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull20895(): T? = if (size > 1) this[1] else null

fun render20895(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20896")
fun step20896(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
class Counter20897 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
