package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21330 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
data class Pair21331<A, B>(val first: A, val second: B)

fun combine21331(left: Int, right: Int): Pair21331<Int, Int> {
    return Pair21331(left + 96, right + 67)
}
sealed class Outcome21332 {
    data class Success(val payload: String) : Outcome21332()
    data object Failure : Outcome21332()
}

fun describe21332(outcome: Outcome21332): String = when (outcome) {
    is Outcome21332.Success -> outcome.payload
    Outcome21332.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21333(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21334(): T? = if (size > 1) this[1] else null

fun render21334(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21335<A, B>(val first: A, val second: B)

fun combine21335(left: Int, right: Int): Pair21335<Int, Int> {
    return Pair21335(left + 27, right + 50)
}
fun <T> List<T>.secondOrNull21336(): T? = if (size > 1) this[1] else null

fun render21336(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21337 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21337 = "registry-41"
    }
}
