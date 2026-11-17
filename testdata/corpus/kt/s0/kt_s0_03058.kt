package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25090 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25090 = "registry-30"
    }
}
data class Pair25091<A, B>(val first: A, val second: B)

fun combine25091(left: Int, right: Int): Pair25091<Int, Int> {
    return Pair25091(left + 37, right + 15)
}
@JvmName("step25092")
fun step25092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
object Registry25093 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25093 = "registry-51"
    }
}
fun <T, R> List<T>.mapChecked25094(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry25095 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25095 = "registry-65"
    }
}
@JvmName("step25096")
fun step25096(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
