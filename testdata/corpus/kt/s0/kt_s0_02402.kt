package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19730 {
    data class Success(val payload: String) : Outcome19730()
    data object Failure : Outcome19730()
}

fun describe19730(outcome: Outcome19730): String = when (outcome) {
    is Outcome19730.Success -> outcome.payload
    Outcome19730.Failure -> "failed"
}
data class Pair19731<A, B>(val first: A, val second: B)

fun combine19731(left: Int, right: Int): Pair19731<Int, Int> {
    return Pair19731(left + 53, right + 45)
}
@JvmName("step19732")
fun step19732(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
@JvmName("step19733")
fun step19733(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
fun <T, R> List<T>.mapChecked19734(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry19735 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19735 = "registry-81"
    }
}
