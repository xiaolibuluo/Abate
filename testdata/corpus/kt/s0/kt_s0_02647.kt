package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21870 {
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
sealed class Outcome21871 {
    data class Success(val payload: String) : Outcome21871()
    data object Failure : Outcome21871()
}

fun describe21871(outcome: Outcome21871): String = when (outcome) {
    is Outcome21871.Success -> outcome.payload
    Outcome21871.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21872(): T? = if (size > 1) this[1] else null

fun render21872(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step21873")
fun step21873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
object Registry21874 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21874 = "registry-43"
    }
}
data class Pair21875<A, B>(val first: A, val second: B)

fun combine21875(left: Int, right: Int): Pair21875<Int, Int> {
    return Pair21875(left + 50, right + 62)
}
fun <T, R> List<T>.mapChecked21876(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter21877 {
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
