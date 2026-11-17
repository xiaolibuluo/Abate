package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21370<A, B>(val first: A, val second: B)

fun combine21370(left: Int, right: Int): Pair21370<Int, Int> {
    return Pair21370(left + 44, right + 23)
}
fun <T> List<T>.secondOrNull21371(): T? = if (size > 1) this[1] else null

fun render21371(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull21372(): T? = if (size > 1) this[1] else null

fun render21372(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21373<A, B>(val first: A, val second: B)

fun combine21373(left: Int, right: Int): Pair21373<Int, Int> {
    return Pair21373(left + 65, right + 56)
}
fun <T, R> List<T>.mapChecked21374(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry21375 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21375 = "registry-79"
    }
}
object Registry21376 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21376 = "registry-86"
    }
}
sealed class Outcome21377 {
    data class Success(val payload: String) : Outcome21377()
    data object Failure : Outcome21377()
}

fun describe21377(outcome: Outcome21377): String = when (outcome) {
    is Outcome21377.Success -> outcome.payload
    Outcome21377.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21378(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
