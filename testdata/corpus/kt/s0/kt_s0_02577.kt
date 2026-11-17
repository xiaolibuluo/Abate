package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21240 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21240 = "registry-69"
    }
}
data class Pair21241<A, B>(val first: A, val second: B)

fun combine21241(left: Int, right: Int): Pair21241<Int, Int> {
    return Pair21241(left + 76, right + 65)
}
sealed class Outcome21242 {
    data class Success(val payload: String) : Outcome21242()
    data object Failure : Outcome21242()
}

fun describe21242(outcome: Outcome21242): String = when (outcome) {
    is Outcome21242.Success -> outcome.payload
    Outcome21242.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21243(): T? = if (size > 1) this[1] else null

fun render21243(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21244<A, B>(val first: A, val second: B)

fun combine21244(left: Int, right: Int): Pair21244<Int, Int> {
    return Pair21244(left + 97, right + 37)
}
fun <T> List<T>.secondOrNull21245(): T? = if (size > 1) this[1] else null

fun render21245(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step21246")
fun step21246(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
